import {
  DynamoDBClient,
  ListTablesCommand,
  paginateListTables
} from '@aws-sdk/client-dynamodb'

class ServeDB {
  constructor () {
    this.dynamoClient = new DynamoDBClient({ region: 'us-east-1' })
    this.listing = new ListTablesCommand()
  }

  get dbClient () {
    return this.dynamoClient
  }

  async tableRequest () {
    try {
      results = await this.dynamoClient.send(this.listing)
    } catch (e) {
      console.error(e)
    }
  }

  page () {
  /*   paginatorConfig = {
      client: this.dynamoClient,
      pageSize: 31
    }

    commandParams = {}
    tableNames = []
    paginator = paginateListTables(paginatorConfig, commandParams)
    
    for await (page of paginator) {
      tableNames.push(...page.TableNames)
    } */ // simplified:
    tableNames = []
    for await (page of paginateListTables(this.dynamoClient, {})) {
      tableNames.push(...page.TableNames)
    }
  }
}
