import './globals.css';

export const metadata = {
  title: 'yarcee',
  description: 'yet another remote code execution engine',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-stone-900	text-white">{children}</body>
    </html>
  );
}
