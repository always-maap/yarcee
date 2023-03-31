import { InputHTMLAttributes } from 'react';
import clsx from 'clsx';

type Props = {} & InputHTMLAttributes<HTMLInputElement>;

export default function TextInput(props: Props) {
  return <input className={clsx(props.className, 'block w-full p-2.5 bg-zinc-800')} />;
}
