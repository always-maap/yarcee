import { ButtonHTMLAttributes } from 'react';

type Props = {} & ButtonHTMLAttributes<HTMLButtonElement>;

export default function Button(props: Props) {
  return (
    <button className="btn-primary" {...props}>
      {props.children}
    </button>
  );
}
