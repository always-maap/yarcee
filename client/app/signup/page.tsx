'use client';

import YARCEE from '@/components/yarcee';
import Button from '@/components/ui/button';
import Label from '@/components/ui/label';
import TextInput from '@/components/ui/text-input';
import Link from 'next/link';
import { SubmitHandler, useForm } from 'react-hook-form';
import { signUp } from '@/api/auth/sign-up';

type Inputs = {
  name: string;
  username: string;
  password: string;
};

export default function SignUp() {
  const { register, handleSubmit } = useForm<Inputs>();

  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    await signUp({ name: data.name, username: data.username, password: data.password });
  };

  return (
    <div className="h-full flex flex-col justify-between items-center py-12">
      <Link href="/">
        <YARCEE width={48} height={48} />
      </Link>
      <form className="flex flex-col space-y-8" onSubmit={handleSubmit(onSubmit)}>
        <div>
          <h1 className="text-center max-w-sm text-6xl mb-4">Sign up for your YARCEE</h1>
          <span>Login or register to start building your projects today.</span>
        </div>
        <div>
          <Label htmlFor="name">Name</Label>
          <TextInput {...register('name')} id="name" placeholder="Mr.example" required />
          <Label htmlFor="username" className="mt-4">
            Username
          </Label>
          <TextInput {...register('username')} id="username" placeholder="example" required />
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
