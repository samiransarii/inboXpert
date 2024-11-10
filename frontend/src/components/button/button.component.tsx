import { ReactNode, FC } from "react";
import "./button.styles.scss";

interface CustomButtonProps {
  children: ReactNode;
  type?: "work" | "personal" | "promotional" | "spam" | "social" | "all-mails";
  handleClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
}

const CustomButton: FC<CustomButtonProps> = ({
  children,
  type,
  handleClick,
}) => {
  const baseClass = "px-6 py-3 rounded rounded-full";
  const typeClass =
    type === "work"
      ? "bg-clr-work text-white"
      : type === "personal"
      ? "bg-clr-personal text-white"
      : type === "promotional"
      ? "bg-clr-promotional text-white"
      : type === "spam"
      ? "bg-clr-spam text-white"
      : type === "social"
      ? "bg-clr-social text-white"
      : type === "all-mails"
      ? "bg-clr-all-mails text-white"
      : "bg-clr-default text-black";

  return (
    <button onClick={handleClick} className={`${baseClass} ${typeClass}`}>
      {children}
    </button>
  );
};

export default CustomButton;
