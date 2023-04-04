import YARCEE from '@/components/yarcee';
import Link from 'next/link';

type Props = {
  name: string;
};

export default function Header(props: Props) {
  return (
    <header className="p-4" style={{ height: 'var(--nav-height)' }}>
      <div className="flex justify-between items-center">
        <Link href="/">
          <YARCEE />
        </Link>
        <div className="text-neutral-400 text-sm">{props.name}</div>
        <Link href="signin">Sign out</Link>
      </div>
    </header>
  );
}
