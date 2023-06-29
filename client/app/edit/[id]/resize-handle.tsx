import { PanelResizeHandle } from 'react-resizable-panels';
import clsx from 'clsx';
import styles from './resize-handle.module.css';

export default function ResizeHandle() {
  return <PanelResizeHandle className={clsx(styles.ResizeHandleOuter, 'w-full h-full rounded')} />;
}
