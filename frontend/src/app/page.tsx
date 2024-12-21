'use client';
import { Button } from '@nextui-org/button';
import { useLaunchParams } from '@telegram-apps/sdk-react';
import Link from 'next/link';

export default function Home() {
  const { initDataRaw } = useLaunchParams();
  console.log(initDataRaw);

  return (
    <div className='h-contentFlow w-full'>
      <Button as={Link} href='/profile' className='h-[50px] w-[100px] bg-uiLightBLueGradient text-uiDeepLightBlue'>
        Button
      </Button>
    </div>
  );
}
