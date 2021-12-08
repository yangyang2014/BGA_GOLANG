## io的file操作注意点
- 需要做close
    比如gzip的压缩，必须使用w.close,如果不使用，直接将res写入文件中。在解压缩软件中无法打开，使用解压缩代码也会报错io.ErrUnexpectedEOF
          var b bytes.Buffer
          w := gzip.NewWriter(&b)
          defer w.Close()
          w.Write(data)
          w.Flush()
          w.Close() //会写入文件结尾符
          res = b.Bytes()