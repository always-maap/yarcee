'use client';

import { Panel, PanelGroup, PanelResizeHandle } from 'react-resizable-panels';
import Header from './header';
import ResizeHandle from './resize-handle';
import Editor from '@uiw/react-codemirror';
import { useCallback, useRef } from 'react';

type Props = {
  params: { id: string };
};

export default function ProjectId({ params }: Props) {
  const editorValue = useRef('');

  const onChange = useCallback((value: string) => {
    editorValue.current = value;
  }, []);

  function onRun() {
    console.log(editorValue.current);
  }

  return (
    <>
      <Header name="test" />
      <PanelGroup direction="horizontal" className="px-4 my-2 gap-2 grow">
        <Panel defaultSize={20}>
          <Editor value="" onChange={onChange} height="500px" theme="dark" spellCheck={false} />
        </Panel>
        <ResizeHandle />
        <Panel className="px-4 bg-black rounded" defaultSize={20}>
          <div className="flex justify-between items-center my-2">
            <span className="text-neutral-400 text-xs">&gt;_ STDOUT</span>
            <button onClick={onRun} className="btn-primary py-0.5 px-2">
              Run
            </button>
          </div>
        </Panel>
      </PanelGroup>
    </>
  );
}
