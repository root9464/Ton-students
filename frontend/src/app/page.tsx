'use client';
import { useBackButton } from '@/shared/hooks/useBackButton';
import { Button } from '@nextui-org/button';
import { Input } from '@nextui-org/input';

export default function Home() {
  useBackButton(false);

  return (
    <div className='h-contentFlow w-full bg-lime-300'>
      <Input label='Email' type='email' />

      <Button color='primary' className='h-[50px] w-[100px] text-white'>
        Button
      </Button>
    </div>
  );
}
