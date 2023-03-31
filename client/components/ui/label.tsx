import clsx from 'clsx';
import { LabelHTMLAttributes } from 'react';

type Props = {} & LabelHTMLAttributes<HTMLLabelElement>;

export default function Label(props: Props) {
  return (
    <label className={clsx(props.className, 'block mb-2 text-sm font-medium text-white')}>
      {props.children}
    </label>
  );
}
