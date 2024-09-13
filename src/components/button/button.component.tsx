import { ReactNode, FC } from "react";
import "./button.styles.scss";

interface CustomButtonProps {
  children: ReactNode;
  type?: "work" | "personal" | "promotional" | "spam" | "social" | "all-mails";
}

const CustomButton: FC<CustomButtonProps> = ({ children, type }) => {
  const baseClass = "px-6 py-3 rounded rounded-full";
  const typeClass =
    type === "work"
      ? "color-work text-white"
      : type === "personal"
      ? "color-personal text-white"
      : type === "promotional"
      ? "color-promotional text-white"
      : type === "spam"
      ? "color-spam text-white"
      : type === "social"
      ? "color-social text-white"
      : type === "all-mails"
      ? "color-all-mails text-white"
      : "color-default text-black";

  return <button className={`${baseClass} ${typeClass}`}>{children}</button>;
};

export default CustomButton;
