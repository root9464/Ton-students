import { NextUIProvider } from '@nextui-org/system';

export const Provider = ({ children }: { children: React.ReactNode }) => {
  return <NextUIProvider>{children}</NextUIProvider>;
};
