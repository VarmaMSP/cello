import Koa from 'koa'
import Router from 'koa-router'
import next from 'next'

import json from 'koa-json'
import bodyParser from 'koa-bodyparser'
import { ParsedUrlQuery } from 'querystring'

const app = next({ dev: process.env.NODE_ENV !== 'production' })
const router = new Router()

router.get('/podcast/:podcastId', async (ctx) => {
  const body: ParsedUrlQuery = { podcast: { id: 'fmasklnfm' } } as any
  await app.render(ctx.req, ctx.res, '/', body)

  ctx.respond = false
})

const server = new Koa()

server.use(json())
server.use(bodyParser())
server.use(router.routes())
server.use(router.allowedMethods())

app.prepare().then(() => {
  const port = parseInt(process.env.PORT || '3000', 10)
  server.listen(port, () => {
    console.log('started')
  })
})
