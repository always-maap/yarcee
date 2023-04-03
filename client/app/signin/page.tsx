import YARCEE from '@/components/yarcee';
import Button from '@/components/ui/button';
import Label from '@/components/ui/label';
import TextInput from '@/components/ui/text-input';
import Link from 'next/link';

export default function SignIn() {
  return (
    <div className="h-full flex flex-col justify-between items-center py-12">
      <Link href="/">
        <YARCEE width={48} height={48} />
      </Link>
      <div className="flex flex-col space-y-8">
        <div>
          <h1 className="text-center max-w-sm text-6xl mb-4">Sign in to your YARCEE</h1>
          <span>Login or register to start building your projects today.</span>
        </div>
        <div>
          <Label htmlFor="email">Email</Label>
          <TextInput type="email" id="email" placeholder="example@domain.com" required />
          <Label htmlFor="password" className="mt-4">
            Password
          </Label>
          <TextInput type="password" id="password" placeholder="********" required />
        </div>
        <Button>Sign In</Button>
      </div>
      <span className="text-xs">
        or{' '}
        <Link href="/signup" className="underline">
          create an account
        </Link>
      </span>
    </div>
  );
}
