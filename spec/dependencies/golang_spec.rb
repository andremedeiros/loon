describe 'Dependencies' do
  describe 'Golang', dependency: true, golang: true do
    versions, default = versions_for :golang

    versions.each do |version|
      it "installs v#{version} correctly" do
        test_dep 'golang', version: version, cmd: 'go version', match: version
      end
    end

    it "installs v#{default} as the default" do
      test_dep 'golang', cmd: 'go version', match: default
    end
  end
end
