import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import NavBarFix from "./components/navbar/navbar";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "Loja Encantto",
  description: "Produtos feitos a mão com muita qualidade",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head> <link rel="icon" href="/logo.svg" sizes="any" /> </head>
      <body
        className={`w-auto h-auto min-h-screen relative bg-[#F0E8E1] ${geistSans.variable} ${geistMono.variable} antialiased`}
      >
      </body>
    </html>
  );
}
