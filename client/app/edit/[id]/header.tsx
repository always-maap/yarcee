import YARCEE from '@/components/yarcee';
import Link from 'next/link';
import { FormEvent, useState } from 'react';
import { z } from 'zod';

type Props = {
  name: string;
  onUpdateTitle(name: string): void;
};

const ZTitleForm = z.string();

export default function Header(props: Props) {
  const [isEditing, setIsEditing] = useState(false);

  function onClickTitle() {
    setIsEditing(true);
  }

  function onSubmitTitle(e: FormEvent<HTMLFormElement>) {
    const formData = new FormData(e.currentTarget);
    e.preventDefault();
    const newTitle = ZTitleForm.parse(formData.get('title'));
    if (newTitle !== props.name) {
      props.onUpdateTitle(newTitle);
    }
    setIsEditing(false);
  }

  function onBlurInput() {
    setIsEditing(false);
  }

  return (
    <header className="p-4" style={{ height: 'var(--nav-height)' }}>
      <div className="flex justify-between items-center">
        <Link href="/dashboard">
          <YARCEE />
        </Link>
        {isEditing ? (
          <form name="title-form" onSubmit={onSubmitTitle}>
            <input
              autoFocus
              name="title"
              onBlur={onBlurInput}
              className="bg-transparent"
              defaultValue={props.name}
            />
          </form>
        ) : (
          <button onClick={onClickTitle} className="text-neutral-400 text-sm">
            {props.name}
          </button>
        )}
        <Link href="signin">Sign out</Link>
      </div>
    </header>
  );
}
