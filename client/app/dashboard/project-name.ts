import { uniqueNamesGenerator, Config, colors, animals } from 'unique-names-generator';

const config: Config = {
  dictionaries: [colors, animals],
  separator: '-',
};

export function generateUniqueProjectName() {
  return uniqueNamesGenerator(config); // continuous-gray-dragonfly
}
