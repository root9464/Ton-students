'use client';
import { useBackButton } from '@/shared/hooks/useBackButton';
import { Button } from '@nextui-org/button';
import { Input } from '@nextui-org/input';
import { useLaunchParams } from '@telegram-apps/sdk-react';

export default function Home() {
  const { initDataRaw } = useLaunchParams();
  console.log(initDataRaw);

  useBackButton(false);

  return (
    <div className='h-contentFlow w-full'>
      <Input label='Email' type='email' />

      <Button color='primary' className='h-[50px] w-[100px] text-white'>
        Button
      </Button>
    </div>
  );
}
