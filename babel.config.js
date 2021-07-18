module.exports = api => {
  api.cache(false);

  const plugins = [
    [
      require.resolve('babel-plugin-module-resolver'),
      {
        extensions: ['.ts', '.tsx', '.png'],
        alias: {
          services: './src/services',
          features: './src/features',
          components: './src/components',
          types: './src/types',
          assets: './assets',
        },
      },
    ],
  ];

  return {
    presets: ['module:metro-react-native-babel-preset'],
    plugins: plugins,
  };
};
