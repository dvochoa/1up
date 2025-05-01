import { ThemeProvider } from "next-themes";
import ThemeToggle from "./theme-toggle";

import styles from "./header.module.css";

const Header = ({ className = "" }: { className?: string }) => {
  return (
    <div className={`${styles["grid-container"]}`}>
      <ThemeProvider attribute="class">
        <ThemeToggle className={`${className} place-self-center`} />
      </ThemeProvider>
    </div>
  );
};

export default Header;
