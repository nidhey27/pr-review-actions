import StringModel from './StringModel';

const yourString = "!dist/**,!**/*.pb.go,!**/*.lock,!**/*.yaml,!**/*.yml,!**/*.cfg,!**/*.toml,!**/*.ini,!**/*.mod,!**/*.sum,!**/*.work,!**/*.json,!**/*.mmd,!**/*.svg,!**/*.png,!**/*.dot,!**/*.md5sum,!**/*.wasm,!**/gen/**,!**/_gen/**,!**/generated/**,!**/vendor/**";

(async () => {
  await sequelize.sync(); // Ensure the table exists

  const newString = await StringModel.create({
    stringValue: yourString,
  });

  console.log('New string created:', newString.stringValue);
})();
