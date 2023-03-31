import './globals.css';

export const metadata = {
  title: 'yarcee',
  description: 'yet another remote code execution engine',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html className="h-full" lang="en">
      <body className="h-full bg-neutral-900	text-white">{children}</body>
    </html>
  );
}
