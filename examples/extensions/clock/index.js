// clock — a ncode extension written in plain Node (no dependencies).
//
// Registers two slash commands:
//   /now            — pushes the current local time into the chat as
//                     a one-shot note (no model call, no transcript)
//   /uptime         — submits a prompt asking the agent to comment on
//                     how long this extension has been running
//
// Why .js and not .ts: this file uses JSDoc types so it can be
// type-checked by tsc / tsserver without a build step. Renaming to
// .ts and updating extension.json's args to ["--import","tsx",
// "index.ts"] (with tsx installed) works too. The extension protocol
// itself is language-agnostic; what matters is that `exec` produces
// a process that reads JSON lines from stdin and writes them to
// stdout.
//
// Install:
//   ncode ext install /path/to/this/dir
//
// Then in ncode:
//   /now
//   /uptime

import { createInterface } from "node:readline";
import { stdin, stdout, stderr } from "node:process";

const NAME = "clock";
const VERSION = "1.0.0";
const STARTED_AT = Date.now();
const ACK_KEYS = new Set([
  "type",
  "product",
  "protocol_version",
  "ncode_version",
  "provider",
  "model",
  "cwd",
  "extension_dir",
  "data_dir",
]);

/** @typedef {{type: string, id?: string, [k: string]: unknown}} Frame */

/**
 * Send a frame to ncode. One JSON object per line; flush immediately
 * so the host doesn't sit waiting on a buffer.
 * @param {Frame} obj
 */
function send(obj) {
  stdout.write(JSON.stringify(obj) + "\n");
}

/**
 * stderr is captured by ncode to $NCODE_HOME/logs/ext-clock.log; perfect
 * for debug output. Anything written to stdout would corrupt the
 * protocol stream.
 * @param {string} msg
 */
function log(msg) {
  stderr.write(`[${NAME}] ${msg}\n`);
}

// 1. Hello first.
send({
  type: "hello",
  name: NAME,
  version: VERSION,
  capabilities: ["commands"],
});

// 2. Wait for the exact ncode protocol-v2 acknowledgement, then register.
function registerCommands() {
  send({
    type: "register_command",
    name: "now",
    description: "show the current local time (no model call)",
  });
  send({
    type: "register_command",
    name: "uptime",
    description: "ask the agent to riff on how long the clock ext has run",
  });
  send({ type: "ready" });
}

// 3. Read frames until stdin closes (ncode shuts us down).
const rl = createInterface({ input: stdin, crlfDelay: Infinity });

rl.on("line", (line) => {
  /** @type {Frame} */
  let frame;
  try {
    frame = JSON.parse(line);
  } catch (err) {
    log(`malformed frame: ${err}`);
    return;
  }

  switch (frame.type) {
    case "hello_ack":
      if (
        frame.product !== "ncode" ||
        frame.protocol_version !== 2 ||
        typeof frame.ncode_version !== "string" ||
        frame.ncode_version === "" ||
        Object.keys(frame).some((key) => !ACK_KEYS.has(key))
      ) {
        log("incompatible host acknowledgement; ncode protocol v2 required");
        process.exit(1);
      }
      log(
        `connected to ncode ${frame.ncode_version} (${frame.provider}/${frame.model})`,
      );
      registerCommands();
      break;

    case "command_invoked":
      handleCommand(frame);
      break;

    case "shutdown":
      send({ type: "shutdown_ack" });
      rl.close();
      break;

    default:
      log(`unknown frame type: ${frame.type}`);
  }
});

rl.on("close", () => {
  log("read loop closed; exiting");
  process.exit(0);
});

/**
 * @param {Frame & {name?: string, args?: string}} frame
 */
function handleCommand(frame) {
  const name = String(frame.name ?? "");
  const args = String(frame.args ?? "").trim();
  const id = String(frame.id ?? "");

  switch (name) {
    case "now": {
      const now = new Date();
      const human = now.toLocaleString(undefined, {
        weekday: "short",
        year: "numeric",
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      });
      const iso = now.toISOString();
      send({
        type: "command_response",
        id,
        action: "display",
        display: `local: ${human}\niso  : ${iso}`,
      });
      return;
    }

    case "uptime": {
      const ms = Date.now() - STARTED_AT;
      const seconds = Math.round(ms / 1000);
      const focus = args ? `Focus on the topic: ${args}.` : "";
      send({
        type: "command_response",
        id,
        action: "prompt",
        prompt:
          `The clock extension has been running for ${seconds}s in this ncode session. ` +
          `Riff on that briefly in one short sentence — be a little dramatic. ${focus}`.trim(),
      });
      return;
    }

    default:
      send({
        type: "command_response",
        id,
        action: "noop",
        error: `clock: unknown command /${name}`,
      });
  }
}
