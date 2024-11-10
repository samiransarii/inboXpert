chrome.runtime.onInstalled.addListener(() => {
  console.log("Extension installed.");
});

chrome.runtime.onMessage.addListener((message, _sender, sendResponse) => {
  if (message.action === "test") {
    console.log("Message received in background script.");
    sendResponse({ success: true });
  }
});
