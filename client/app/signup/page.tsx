'use client';

import Logo from '@/components/logo';
import Button from '@/components/ui/button';
import Label from '@/components/ui/label';
import TextInput from '@/components/ui/text-input';
import Link from 'next/link';
import { SubmitHandler, useForm } from 'react-hook-form';

type Inputs = {
  email: string;
  password: string;
};

export default function SignUp() {
  const { register, handleSubmit } = useForm<Inputs>();

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    console.log(data);
  };

  return (
    <div className="h-full flex flex-col justify-between items-center py-12">
      <Link href="/">
        <Logo width={48} height={48} />
      </Link>
      <form className="flex flex-col space-y-8" onSubmit={handleSubmit(onSubmit)}>
        <div>
          <h1 className="text-center max-w-sm text-6xl mb-4">Sign up for your YARCEE</h1>
          <span>Login or register to start building your projects today.</span>
        </div>
        <div>
          <Label htmlFor="email">Email</Label>
          <TextInput
            {...register('email')}
            type="email"
            id="email"
            placeholder="example@domain.com"
            required
          />
          <Label htmlFor="password" className="mt-4">
            Password
          </Label>
          <TextInput
            {...register('password')}
            type="password"
            id="password"
            placeholder="********"
            required
          />
        </div>
        <Button>Sign Up</Button>
      </form>
      <span className="text-xs">
        or{' '}
        <Link href="/signin" className="underline">
          sign in to your account
        </Link>
      </span>
    </div>
  );
}
