'use client';

import { Panel, PanelGroup, PanelResizeHandle } from 'react-resizable-panels';
import Header from './header';
import ResizeHandle from './resize-handle';
import Editor from '@uiw/react-codemirror';
import { useCallback, useRef } from 'react';
import { useSandbox } from '@/hooks/useSandbox';
import { updateSandbox } from '@/api/sandbox/update-sandbox';
import { mutate } from 'swr';
import { executeSandbox } from '@/api/sandbox/execute-sandbox';

type Props = {
  params: { id: string };
};

export default function ProjectId({ params }: Props) {
  const { sandbox, isLoading } = useSandbox(params.id);
  const editorValue = useRef('');

  const onChange = useCallback((value: string) => {
    editorValue.current = value;
  }, []);

  if (isLoading) {
    return 'loading...';
  }

  async function onRun() {
    console.log(editorValue.current);
    const resp = await executeSandbox({ id: params.id, code: editorValue.current });
    console.log(resp);
  }

  async function onUpdateTitle(newTitle: string) {
    const resp = await updateSandbox({
      id: params.id,
      name: newTitle,
      language: sandbox!.language,
      code: sandbox!.code,
    });
    mutate(`/sandbox/${params.id}`);
  }

  return (
    <>
      <Header name={sandbox!.name} onUpdateTitle={onUpdateTitle} />
      <PanelGroup direction="horizontal" className="px-4 my-2 gap-2 grow">
        <Panel defaultSize={50}>
          <Editor
            value={sandbox!.code}
            onChange={onChange}
            height="100%"
            style={{ height: '100%' }}
            theme="dark"
            spellCheck={false}
          />
        </Panel>
        <div className="w-1">
          <ResizeHandle />
        </div>
        <Panel>
          <PanelGroup direction="vertical" className="gap-2">
            <Panel className="px-4 bg-black rounded" defaultSize={70}>
              <div className="flex justify-between items-center my-2">
                <span className="text-neutral-400 text-xs">&gt;_ STDOUT</span>
                <button onClick={onRun} className="btn-primary py-0.5 px-2">
                  Run
                </button>
              </div>
              <div className="whitespace-pre-wrap">{sandbox?.stdout}</div>
            </Panel>
            <div className="h-1">
              <ResizeHandle />
            </div>
            <Panel className="px-4 bg-black rounded">
              <div className="flex justify-between items-center my-2">
                <span className="text-neutral-400 text-xs">&gt;_ STDERR</span>
              </div>
              {sandbox?.stderr}
            </Panel>
          </PanelGroup>
        </Panel>
      </PanelGroup>
    </>
  );
}
