import { ReactNode } from 'react';

type Props = {
  icon: ReactNode;
  actions?: ReactNode;
  name: ReactNode;
  info?: ReactNode;
};

export default function ProjectCard(props: Props) {
  return (
    <div className="bg-neutral-800 hover:bg-neutral-700 p-6 flex flex-col justify-between h-full rounded">
      <div className="flex justify-between items-center">
        {props.icon} {props.actions}
      </div>
      <div>
        {props.name} {props.info}
      </div>
    </div>
  );
}
