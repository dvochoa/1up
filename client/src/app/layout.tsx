import { ThemeProvider } from "next-themes";

import "@/styles/global.css";

export const metadata = {
  title: "1up",
  description: "A 10000 hours productivy app",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <ThemeProvider attribute="class">{children}</ThemeProvider>
      </body>
    </html>
  );
}
