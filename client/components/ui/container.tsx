import { HTMLAttributes } from 'react';
import clsx from 'clsx';

type Props = {} & HTMLAttributes<HTMLDivElement>;

export default function Container(props: Props) {
  return (
    <div {...props} className={(props.className, 'mx-auto max-w-[1234px]')}>
      {props.children}
    </div>
  );
}
