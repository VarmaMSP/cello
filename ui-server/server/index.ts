import Koa from 'koa'
import Router from 'koa-router'
import next from 'next'
import bodyParser from 'koa-bodyparser'
import { ParsedUrlQuery } from 'querystring'

const app = next({ dev: process.env.NODE_ENV !== 'production' })
const handle = app.getRequestHandler()
const router = new Router()

router.get('/podcast/:podcastId', async (ctx) => {
  const body: ParsedUrlQuery = { podcast: { id: 'fmasklnfm' } } as any
  await app.render(ctx.req, ctx.res, '/', body)
  ctx.respond = false
})

router.get('*', async (ctx) => {
  await handle(ctx.req, ctx.res)
  ctx.respond = false
})

const server = new Koa()

server.use(bodyParser())
server.use(router.routes())
server.use(router.allowedMethods())

app.prepare().then(() => {
  const port = parseInt(process.env.PORT || '3000', 10)
  server.listen(port, () => {
    console.log('UI server running on port 3000')
  })
})
