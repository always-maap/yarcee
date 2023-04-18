import Image from 'next/image';

import Nodejs from '@/components/nodejs';

export const TEMPLATES = [
  { name: 'Node.js', icon: <Nodejs />, code: 'console.log("called")' },
  {
    name: 'Python',
    icon: <Image width={20} height={20} src="/logo/python.png" alt="python-logo" />,
    code: 'print("hi")',
  },
];
