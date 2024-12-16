// @ts-check
const path = require("path");

/**
 * Returns path to ticketeer binary for current platform
 * @returns {string}
 */
function getBinaryPath() {
  /**
   * @type {string}
   */
  let os = process.platform;
  let extension = "";
  if (os === "win32" || os === "cygwin") {
    os = "windows";
    extension = ".exe";
  }
  const arch = process.arch;

  const binaryPackagePath = path.join(
    `ticketeer-${os}-${arch}`,
    `package.json`,
  );
  const binaryPackageDir = path.dirname(require.resolve(binaryPackagePath));
  const binaryPath = path.join(binaryPackageDir, `ticketeer${extension}`);

  return binaryPath;
}

exports.getBinaryPath = getBinaryPath;
