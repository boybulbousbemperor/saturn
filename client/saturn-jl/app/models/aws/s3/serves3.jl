using AWS.AWSServices: s3
using AWS: @service

struct ServeS3
  s3Client::@service S3

  itemKey::ItemParameter
  bucketName::String
  htmlBody::String
  httpFunc::String
  ssoPath::String

  function getKey()
    return itemKey
  end

  function getBucketName()
    return bucketName
  end

  function getHTMLBody()
    return htmlBody
  end

  function getHTTPMethod()
    return httpFunc
  end

  function getSSOPath()
    return ssoPath
  end

  function setKey(key)
    ServeS3.::itemKey = key
  end

  function setBucketName(bucket)
    ServeS3.::bucketName = bucket
  end

  function setHTMLBody(body)
    ServeS3.::htmlBody = body
  end

  function setHTTPMethod(method)
    ServeS3.::httpFunc = method
  end

  function setSSOPath(path)
    ServeS3.::ssoPath = path
  end

  function retreiveItem(httpFunc, bucketName)
    s3Client(httpFunc, bucketName)
  end

  function body(getSSOPath(), getKey())
    return """
    <Get xmlns="$getSSOPath">
        <Object>
            <Key>$getKey</Key>
        </Object>
    </Get>
    """
  end

  function testRetreival()
    retreiveItem("GET", body(bucketName, htmlBody))
  end
end

struct ItemParameter
  keyID::String
  keyValue::String
end

@service S3
s3("GET", "/ds-infra-serv")
S3.list_objects("/ds-infra-serv")
