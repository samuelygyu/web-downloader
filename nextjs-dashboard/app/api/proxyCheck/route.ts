import axios, { AxiosResponse } from 'axios'
import { log } from 'console';
import { NextRequest, NextResponse } from 'next/server';

export async function GET(req: NextRequest) {
  const searchParams = req.nextUrl.searchParams;
  const host = searchParams.get('host') as string;
  const port = parseInt(searchParams.get('port') as string);

  console.log(host, port)
  let status = false;
  await axios.get('https://www.google.com', {
    proxy: {
      protocol: 'http',
      host: host,
      port: port
    }
  }).then(res => {
    status = res.status === 200;
  }).catch(err => {
    status = false;
  })

  return NextResponse.json({ status });
}