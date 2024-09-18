console.log("Chrome extension running!");

chrome.runtime.onInstalled.addListener(() => {
  console.log("Extension installed.");
});
