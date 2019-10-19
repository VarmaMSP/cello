import ButtonPlay from 'components/button_play'
import EpisodeMeta from 'components/episode_meta'
import Grid from 'components/grid'
import Link from 'next/link'
import { Episode } from 'types/app'
import { getImageUrl } from 'utils/dom'

export interface StateToProps {
  feed: Episode[]
}

export interface DispatchToProps {
  showEpisodeModal: (episodeId: string) => void
}

interface Props extends StateToProps, DispatchToProps {}

const ListFeed: React.SFC<Props> = (props) => {
  const { feed, showEpisodeModal } = props

  return (
    <>
      <h1 className="text-xl text-gray-900 mb-5">
        {'Your Feed for last week'}
      </h1>
      <Grid
        cols={{ LG: 3, MD: 1, SM: 1 }}
        classNameChild="flex my-2 lg:px-2 py-2 rounded-lg md:hover:bg-gray-200"
        totalRowSpacing={{ LG: 2, MD: 10, SM: 0 }}
      >
        {feed.map((episode) => (
          <>
            <Link
              href={{
                pathname: '/podcasts',
                query: { podcastId: episode.podcastId },
              }}
              as={`/podcasts/${episode.podcastId}`}
              key={episode.podcastId}
            >
              <a>
                <img
                  className="w-24 h-24 flex-none object-contain rounded-lg border cursor-pointer"
                  src={getImageUrl(episode.podcastId, 'md')}
                />
              </a>
            </Link>
            <div className="flex flex-col justify-between pl-3">
              <div>
                <h1
                  className="text-sm leading-tight line-clamp-2"
                  onClick={() => showEpisodeModal(episode.id)}
                >
                  {episode.title}
                </h1>
                <div className="mt-2">
                  <EpisodeMeta episode={episode} />
                </div>
              </div>
              <ButtonPlay className="w-5" episodeId={episode.id} />
            </div>
          </>
        ))}
      </Grid>
    </>
  )
}

export default ListFeed
