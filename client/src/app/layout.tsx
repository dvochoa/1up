import { ThemeProvider } from "next-themes";
import { Merriweather } from "next/font/google";

import ThemeToggle from "@/components/theme-toggle";
import styles from "./layout.module.css";

const merriweather = Merriweather({
  weight: ["300", "400", "700", "900"],
  style: ["normal", "italic"],
  subsets: ["latin"],
  display: "swap",
});

import "@/styles/global.css";

export const metadata = {
  title: "1up",
  description: "A 10000 hours productivy app",
};

// TODO: Move footer out into its own component
function Footer() {
  return <div className="border-2 border-solid">Hello</div>;
}

// TODO: Add a header component and make themeToggle a part of that
// TODO: How to enforce that
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className={merriweather.className} suppressHydrationWarning>
      <body className={`${styles["grid-container"]} grid h-lvh`}>
        <ThemeProvider attribute="class">
          <ThemeToggle />
        </ThemeProvider>
        {children}
        <Footer />
      </body>
    </html>
  );
}
