import "@/styles/global.css";

export const metadata = {
  title: "dannyochoa",
  description: "Daniel Ochoa's personal website",
};
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
