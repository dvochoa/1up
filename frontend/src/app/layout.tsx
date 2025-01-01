import "@/styles/global.css";

export const metadata = {
  title: "1up",
  description: "A 10000 hours productivy app",
};
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
