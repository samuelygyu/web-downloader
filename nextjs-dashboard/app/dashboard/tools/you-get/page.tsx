'use client'

import axios from "axios"
import { useState } from "react"
import { useForm, SubmitHandler } from "react-hook-form"
import { FcApproval, FcDisclaimer } from "react-icons/fc";
import styles from '@/app/ui/home.module.css'
import Link from "next/link";

type DownloadForm = {
  url: string,
  proxyHost: string,
  proxyPort: string,
  cookies: FileList
}

axios.interceptors.request.use(request => {
  console.log(request.data);
  return request;
})

export default function Page() {
  const { register, handleSubmit, getValues } = useForm<DownloadForm>();
  const [status, setStatus] = useState<boolean>(false)
  const [isVisible, setIsVisible] = useState<boolean>(false)
  const [files, setFiles] = useState<string[]>([])

  async function handleProxyCheck() {
    let host = getValues('proxyHost')
    let port = getValues('proxyPort')
    if (!host || !port) {
      alert("请输入代理地址")
      return
    }
    axios.get('http://localhost:3000/api/proxyCheck', {
      params: {
        host: getValues('proxyHost'),
        port: getValues('proxyPort')
      }
    }).then(res => {
      setStatus(res.data.status)
      setIsVisible(true)
    }).catch(err => {
      console.error(err)
    })
  }

  const onSubmit: SubmitHandler<DownloadForm> = (formData) => {
    const form = new FormData()
    console.log(formData.cookies)

    if (formData.cookies !== null) {
      form.append('cookies', formData.cookies[0])
    }
    if (formData.proxyHost && formData.proxyPort) {
      form.append('proxy', formData.proxyHost + ":" + formData.proxyPort)
    }
    form.append('url', formData.url);

    axios.post('http://175.178.54.54:8080/parse', form, {
      headers: {
        'Content-Type':'multipart/form-data'
      }
    }).then(res => {
      console.log(res)
      setFiles(res.data.files)
    }).catch(err => {
      console.error(err)
    })
  }
    
  return (
    <>
      <h1 className="text-2xl font-bold underline">视频下载器</h1>

      <form onSubmit={handleSubmit(onSubmit)}>
        <div>
          <input {...register("url", { required: true})} type="text" placeholder="url"/>
          <label>下载链接地址</label>
        </div>
        <div className={styles.hostPortContainer}>
          <input {...register("proxyHost")} type="text" placeholder="proxy host" className="host"/>
          <span className="sperator">:</span>
          <input {...register("proxyPort")} type="text" placeholder="proxy port" className="port"/>
          <label style={{marginLeft: '20px'}}>代理地址</label>
          <button className="h-10 px-6 font-semibold rounded-md bg-black text-white" onClick={handleProxyCheck} style={{marginLeft: '20px'}}>测试</button>
          {isVisible && (status ? <FcApproval size={25} className="proxyCheckStatus"/> : <FcDisclaimer size={25} className="proxyCheckStatus"/>)}
        </div>
        <div>
          <input {...register("cookies")} type="file" />
          <button className="text-red">请点击上传cookies文件</button>
        </div>
        <button className="h-10 px-6 font-semibold rounded-md bg-black text-white" type="submit">解析</button>
        <ol>
          <Files fileList={files}/>
        </ol>
      </form>
    </>
  )
}

function Files({fileList}: {fileList: string[]}) {
  function handleClick(fileName: string) {
    axios.post("http://175.178.54.54:8080/download", {
        Filename: fileName
      }, {
      headers: {
        'Content-Type': 'application/json',
      },
      responseType: 'blob'
    }).then(res => {
      console.log(res.data)
      const blob = res.data
      const url = URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = 'reponse'
      document.body.appendChild(a)
      a.click()

      document.body.removeChild(a)
      URL.revokeObjectURL(url)
    })
  }
  return (
    <>
      {fileList.map(file => {
        return (
          <li key={file}>
            <p>{file}</p><button  type="button" onClick={() => handleClick(file)}>下载</button>
          </li>
        )
      })}
    </>
  )
}