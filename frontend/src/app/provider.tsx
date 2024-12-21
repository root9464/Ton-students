'use client';
import { useClientOnce } from '@/shared/hooks/useClientOnce';
import { init } from '@/shared/lib/initTma';
import { ReactNode } from 'react';

export const Provider = ({ children }: Readonly<{ children: ReactNode }>) => {
  useClientOnce(() => {
    init(true);
  });

  return <>{children}</>;
};
