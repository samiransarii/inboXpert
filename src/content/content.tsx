import ReactDOM from "react-dom/client";
import "./content.styles.scss";
import "material-icons/iconfont/material-icons.css";

import TopbarActions from "./top-bar/top-bar.component";

// Utility function to wait for an element to appear in the DOM
function waitForElement(
  selector: string,
  timeout: number = 10000
): Promise<Element> {
  return new Promise((resolve, reject) => {
    const element = document.querySelector(selector);
    if (element) {
      return resolve(element);
    }

    const observer = new MutationObserver((_mutations, obs) => {
      const el = document.querySelector(selector);
      if (el) {
        resolve(el);
        obs.disconnect();
      }
    });

    observer.observe(document.body, { childList: true, subtree: true });

    // Timeout after specified duration
    setTimeout(() => {
      observer.disconnect();
      reject(new Error(`Element ${selector} not found within ${timeout}ms`));
    }, timeout);
  });
}

// Function to fix Gmail's CSS styles
function fixGmailStyles() {
  const selectors = {
    hamBurgerMenu: ".gb_Jc",
    userAvatar: ".aju",
    fixElement: ".nH.aqK",
  };

  const hamBurgerMenu = document.querySelector(selectors.hamBurgerMenu);
  const userAvatar = document.querySelector(selectors.userAvatar);
  const fixElement = document.querySelector(selectors.fixElement);

  if (hamBurgerMenu instanceof HTMLElement) {
    hamBurgerMenu.style.width = "48px";
    hamBurgerMenu.style.height = "48px";
    console.log("Fixed hamburger menu");
  }

  if (userAvatar instanceof HTMLElement) {
    userAvatar.style.minWidth = "80px";
    console.log("Fixed user avatar");
  }

  if (fixElement instanceof HTMLElement) {
    fixElement.style.alignItems = "center";
    console.log("Fixed alignment of fixElement");
  }
}

// Function to inject the React component
function injectReactComponent(targetElement: Element) {
  // Prevent multiple injections
  if (document.getElementById("react-content-script-container")) {
    return;
  }

  const appContainer = document.createElement("div");
  appContainer.id = "react-content-script-container";
  appContainer.style.position = "relative";
  appContainer.style.zIndex = "9999";
  targetElement.insertAdjacentElement("afterend", appContainer);

  const root = ReactDOM.createRoot(appContainer);
  root.render(<TopbarActions />);
}

// Function to handle mutations
function handleMutations(
  mutationsList: MutationRecord[],
  observer: MutationObserver
) {
  console.log(mutationsList);
  for (const mutation of mutationsList) {
    if (mutation.type === "childList" && mutation.addedNodes.length > 0) {
      // Check if the target element is added
      mutation.addedNodes.forEach(async (node) => {
        if (!(node instanceof Element)) return;

        if (node.matches(".aFg")) {
          fixGmailStyles();
          // const targetElement = await waitForElement(".G-tF");
          // injectReactComponent(targetElement);
          observer.disconnect();
        }
      });
    }
  }
}

// Initialize the content script
(async function initializeContentScript() {
  try {
    // Wait for the main container that holds the email view
    const topbarActionsContainer = ".Cq.aqL";
    const emailViewContainer = ".nH";
    const targetElement = await waitForElement(topbarActionsContainer);
    const emailViewContainerElement = await waitForElement(emailViewContainer);

    // Fix styles initially
    fixGmailStyles();

    // Inject the React component
    injectReactComponent(targetElement);

    // Setup MutationObserver to handle dynamic changes
    const observer = new MutationObserver(handleMutations);
    observer.observe(emailViewContainerElement, {
      childList: true,
      subtree: true,
    });

    // Optional: Send a test message to background script
    chrome.runtime.sendMessage({ action: "test" }, (response) => {
      if (chrome.runtime.lastError) {
        console.error("Error sending message:", chrome.runtime.lastError);
      } else {
        console.log("Background response:", response);
      }
    });
  } catch (error) {
    console.error("Content script initialization failed:", error);
  }
})();
