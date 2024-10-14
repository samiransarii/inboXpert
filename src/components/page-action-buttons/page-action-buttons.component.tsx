import { ReactNode, FC } from "react";
import "./page.action.button.styles.scss";

interface PageActionButtonProps {
  children: ReactNode;
  handleClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
  type?: "unsubscribe" | "labels" | "bulkDelete";
}

const PageActionButton: FC<PageActionButtonProps> = ({
  children,
  handleClick,
  type = "unsubscribe",
}) => {
  const baseClass =
    "px-4 py-1.5 rounded rounded-full bg-clr-default hover:bg-clr-default-hover text-black";
  const typeClass =
    type === "unsubscribe"
      ? "unsubscribe"
      : type === "labels"
      ? "labels"
      : type === "bulkDelete"
      ? "bulk-delete"
      : "help";

  function getIcon() {
    switch (type) {
      case "unsubscribe":
        return "unsubscribe";
      case "labels":
        return "label";
      case "bulkDelete":
        return "delete_forever";
      default:
        return "help";
    }
  }

  return (
    <button onClick={handleClick} className={`${baseClass} page-action-button`}>
      <span className={`material-icons-outlined ${typeClass}`}>
        {getIcon()}
      </span>
      {children}
    </button>
  );
};

export default PageActionButton;
