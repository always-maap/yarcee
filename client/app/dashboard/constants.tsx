import Image from 'next/image';

import Nodejs from '@/components/nodejs';

export const TEMPLATES = [
  { name: 'Node.js', icon: <Nodejs />, code: 'console.log("called")', abbr: 'node' },
  {
    name: 'Python',
    icon: <Image width={20} height={20} src="/logo/python.png" alt="python-logo" />,
    code: 'print("hi")',
    abbr: 'py',
  },
  {
    name: 'C',
    icon: <Image width={20} height={20} src="/logo/c.png" alt="c-logo" />,
    code: `#include <stdio.h>
    int main() {
       printf("Hello, World!");
       return 0;
    }`,
    abbr: 'cpp',
  },
  {
    name: 'C++',
    icon: <Image width={20} height={20} src="/logo/cpp.png" alt="cpp-logo" />,
    code: `#include <iostream>

  int main() {
      std::cout << "Hello World!";
      return 0;
  }`,
    abbr: 'cpp',
  },
];
