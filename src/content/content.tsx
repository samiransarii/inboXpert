import ReactDOM from "react-dom/client";
import "./content.styles.scss";
import "material-icons/iconfont/material-icons.css";

import TopbarActions from "./top-bar/top-bar.component";

// Render the component
const appContainer = document.createElement("div");
appContainer.id = "react-content-script-container";
appContainer.style.zIndex = "9999";
document.body.appendChild(appContainer);

function injectButton() {
  const targetElement = document.querySelector(".Cq.aqL");
  const fixElement = document.querySelector(".nH.aqK");

  if (fixElement) {
    (fixElement as HTMLElement).style.alignItems = "center";
  }

  if (targetElement) {
    targetElement.insertAdjacentElement("afterend", appContainer);

    obserber.disconnect();
  } else {
    console.log("Trying to initialize elements.");
  }
}

const obserber = new MutationObserver(() => {
  injectButton();
});

setTimeout(() => {
  obserber.observe(document.body, { childList: true, subtree: true });
  console.log("Started observing for Gmail changes...");
}, 3000);

const root = ReactDOM.createRoot(appContainer);
root.render(
  <>
    <TopbarActions />
  </>
);
