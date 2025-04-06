import { ThemeProvider } from "next-themes";

import { Merriweather } from "next/font/google";

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

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className={merriweather.className} suppressHydrationWarning>
      <body>
        <ThemeProvider attribute="class">{children}</ThemeProvider>
      </body>
    </html>
  );
}
