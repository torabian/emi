import { WebSocketX } from './sdk/common/WebSocketX';
import { YokeMessageDto } from './YokeMessageDto';
import { buildUrl } from './sdk/common/buildUrl';
/**
* Action to communicate with the action streamYoke
*/
	/**
 * StreamYokeAction
 */
export class StreamYokeAction {
  static URL = '/stream-yoke';
  static NewUrl = (
	qs
  ) => buildUrl(
		StreamYokeAction.URL,
		 undefined,
		qs
	);
  static Method = 'reactive';
	static Create = (
		overrideUrl,
		qs,
	) => {
		const url = overrideUrl ?? StreamYokeAction.NewUrl(
			qs
		)
		return new WebSocketX(
			url,
			undefined,
			{
				MessageFactoryClass: undefined,
			}
		);
	}
  static Definition = {
  "name": "streamYoke",
  "url": "/stream-yoke",
  "method": "reactive",
  "description": "Used by the kotlin to golang, to send the yoke information",
  "in": {
    "dto": "YokeMessageDto"
  }
}
}