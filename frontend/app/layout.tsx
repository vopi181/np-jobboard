import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Link from "next/link";
import styles from "@/app/page.module.css";
import { cookies } from "next/headers";

import SignInButton from "@/app/components/SignInButton";
import SignOutButton from "@/app/components/SignOutButton";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Encore + Next.js",
};

const navLinks = [
  { href: "/", label: "Home" },
  // { href: "/users", label: "User List" },
];

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const isLoggedIn = cookies().has("auth-token");

  return (
    <html lang="en">
      <body className={inter.className}>
        <header>
          <nav className={styles.nav}>
            <div className={styles.navLinks}>
              {navLinks.map(({ href, label }) => (
                <Link key={href} href={href}>
                  {label}
                </Link>
              ))}
              
            </div>

            <SignInButton loggedin={isLoggedIn}/>
          </nav>
        </header>

        <main className={styles.main}>{children}</main>
      </body>
    </html>
  );
}
