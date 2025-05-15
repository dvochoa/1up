import { Merriweather } from "next/font/google";

import Header from "./components/header";

const merriweather = Merriweather({
  weight: ["300", "400", "700", "900"],
  style: ["normal", "italic"],
  subsets: ["latin"],
  display: "swap",
});

import "global.css";

export const metadata = {
  title: "1up",
  description: "A 10000 hours productivy app",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className={merriweather.className} suppressHydrationWarning>
      <body className="flex min-h-lvh flex-col">
        <Header className="h-[15vh] flex-shrink-0 lg:h-[10vh]" />
        <main className="flex-grow pb-[5vh]">{children}</main>
      </body>
    </html>
  );
}
