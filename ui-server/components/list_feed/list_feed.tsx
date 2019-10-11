import ButtonWithIcon from 'components/button_with_icon'
import Grid from 'components/grid'
import { NextSeo } from 'next-seo'
import { Episode } from 'types/app'
import { getImageUrl } from 'utils/dom'
import { formatEpisodeDuration, formatEpisodePubDate } from 'utils/format'

export interface StateToProps {
  feed: Episode[]
}

export interface DispatchToProps {
  playEpisode: (episodeId: string) => void
}

interface Props extends StateToProps, DispatchToProps {}

const ListFeed: React.SFC<Props> = (props) => {
  return (
    <>
      <NextSeo
        title="Trending Podcasts - Phenopod"
        description="Trending podcasts"
        canonical="https://phenopod.com/trending"
        openGraph={{
          url: 'https://phenopod.com/trending',
          type: 'article',
          title: 'Trending Podcasts',
          description: 'Trending Podcasts',
          images: [
            { url: getImageUrl(props.feed[2].podcastId, 'sm') },
            { url: getImageUrl(props.feed[3].podcastId, 'sm') },
            { url: getImageUrl(props.feed[5].podcastId, 'sm') },
            { url: getImageUrl(props.feed[7].podcastId, 'sm') },
          ],
        }}
      />
      <h1 className="text-xl text-gray-900 mb-5">
        {'Your Feed for last week'}
      </h1>
      <Grid
        cols={{ LG: 3, MD: 1, SM: 1 }}
        classNameChild="flex my-2 px-2 py-2 rounded-lg md:hover:bg-gray-200"
        totalRowSpacing={{ LG: 2, MD: 10, SM: 0 }}
      >
        {props.feed.map((episode) => (
          <>
            <img
              className="w-24 h-24 flex-none object-contain rounded-lg border cursor-pointer"
              src={getImageUrl(episode.podcastId, 'md')}
            />
            <div className="flex flex-col justify-between pl-3">
              <div>
                <h1 className="text-sm leading-tight line-clamp-2">
                  {episode.title}
                </h1>
                <span className="text-xs text-gray-700">
                  {formatEpisodePubDate(episode.pubDate)}
                  <span className="mx-2 font-extrabold">&middot;</span>
                  {formatEpisodeDuration(episode.duration)}
                </span>
              </div>
              <div className="">
                <ButtonWithIcon
                  className="flex-none w-5 text-gray-700 hover:text-black"
                  icon="play-outline"
                  onClick={() => props.playEpisode(episode.id)}
                />
              </div>
            </div>
          </>
        ))}
      </Grid>
    </>
  )
}

export default ListFeed
