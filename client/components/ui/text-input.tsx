import { ComponentPropsWithRef, forwardRef, InputHTMLAttributes } from 'react';
import clsx from 'clsx';

type Props = {} & InputHTMLAttributes<HTMLInputElement>;

const TextInput = forwardRef<HTMLInputElement, Props>((props, forwardedRef) => {
  return (
    <input
      className={clsx(props.className, 'block w-full p-2.5 bg-zinc-800')}
      ref={forwardedRef}
      {...props}
    />
  );
});

export default TextInput;
