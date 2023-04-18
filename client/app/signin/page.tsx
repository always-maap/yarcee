'use client';

import YARCEE from '@/components/yarcee';
import Button from '@/components/ui/button';
import Label from '@/components/ui/label';
import TextInput from '@/components/ui/text-input';
import Link from 'next/link';
import { SubmitHandler, useForm } from 'react-hook-form';
import { signIn } from '@/api/auth/sign-in';
import Cookies from 'js-cookie';
import { useRouter } from 'next/navigation';

type Inputs = {
  username: string;
  password: string;
};

export default function SignIn() {
  const { register, handleSubmit } = useForm<Inputs>();
  const { push } = useRouter();

  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    const resp = await signIn({ username: data.username, password: data.password });
    Cookies.set('jwt-token', resp.data);
    await push('/dashboard');
  };

  return (
    <div className="h-full flex flex-col justify-between items-center py-12">
      <Link href="/">
        <YARCEE width={48} height={48} />
      </Link>
      <form className="flex flex-col space-y-8" onSubmit={handleSubmit(onSubmit)}>
        <div>
          <h1 className="text-center max-w-sm text-6xl mb-4">Sign in to your YARCEE</h1>
          <span>Login or register to start building your projects today.</span>
        </div>
        <div>
          <Label htmlFor="username">Username</Label>
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
        <Button>Sign In</Button>
      </form>
      <span className="text-xs">
        or{' '}
        <Link href="/signup" className="underline">
          create an account
        </Link>
      </span>
    </div>
  );
}
