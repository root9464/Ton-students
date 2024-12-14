'use client';
import { init, isTMA } from '@telegram-apps/sdk-react';
import { ReactNode } from 'react';

if (await isTMA()) {
  init();
}

export const Provider = ({ children }: Readonly<{ children: ReactNode }>) => {
  return <>{children}</>;
};
