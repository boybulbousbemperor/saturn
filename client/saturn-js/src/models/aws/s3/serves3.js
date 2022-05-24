import { AbortController } from '@aws-sdk/abort-controller'
import {
  CreateBucketCommand,
  GetObjectCommand,
  PutObjectCommand,
  S3Client
} from '@aws-sdk/client-s3'

// Bare-bones clients/commands:
/* const clientS3 = new S3Client()
await clientS3.get(new GetObjectCommand()) */

// Aggregated clients/commands: (less performant than barebones but equivalent)
const serves3 = new S3
await serves3.GetObjectCommand

class ServeS3 {
  constructor () {}

  request () {
    // Abort Controller: AWS SDK for JS
    const abortController = new AbortController
    const requestPromise = serves3.send(new CreateBucketCommand, {
      abortSignal: abortController.signal
    })

    // The abortController can be aborted any time.
    // The request will not be created if abortSignal is already aborted.
    // The request will be destroyed if abortSignal is aborted before response is returned.
    abortController.abort();

    // This will fail with "AbortError" as abortSignal is aborted.
    await requestPromise;
  }

  get s3Item () {
    return this.serves3.GetObjectCommand
  }
  

  set update (serves3) {
    this.serves3.PutObjectCommand
  }
}


