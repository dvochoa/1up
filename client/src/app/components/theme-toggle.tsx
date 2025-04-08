"use client";

import { useState, useEffect } from "react";
import { useTheme } from "next-themes";
import Image from "next/image";

const ICON_WIDTH = 20;
const ICON_HEIGHT = 20;

const ThemeToggle = () => {
  const [mounted, setMounted] = useState(false);
  const { systemTheme, theme, setTheme } = useTheme();
  const currentTheme = theme === "system" ? systemTheme : theme;

  // Determining the theme requires being mounted client side so by rendering only on the client we can avoid hydration errors
  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return null;
  }

  return (
    <button
      onClick={() => (currentTheme == "dark" ? setTheme("light") : setTheme("dark"))}
      className="place-self-center"
    >
      <Image
        src={
          (currentTheme || "light") === "dark"
            ? "/images/dark-mode-icon.svg"
            : "/images/light-mode-icon.svg"
        }
        width={ICON_WIDTH}
        height={ICON_HEIGHT}
        alt="A Sun/Moon that when clicked toggles the color scheme between light and dark mode"
      />
    </button>
  );
};

export default ThemeToggle;
