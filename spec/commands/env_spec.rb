describe 'Commands' do
  describe 'env', command: true do
    it 'should show default values' do
      loon %(env)
      assert_stdout 'PROJECT_IP='
    end

    describe 'without .env' do
      it "should show the project's environment values" do
        with_payload(env: {'LOON' => 'WAS HERE'}) do
          loon %(env)
          assert_stdout 'LOON=WAS HERE'
        end
      end
    end

    describe 'with .env' do
      it "should overwrite the project's environment values" do
        with_payload(env: {'LOON' => 'WAS HERE'}) do |root|
          File.write(File.join(root, '.env'), <<~STR)
            LOON="SOMETHING ELSE"
          STR

          loon %(env)
          assert_stdout 'LOON=SOMETHING ELSE'
        end
      end
      it "should handle all sorts of values" do
        with_payload do |root|
          File.write(File.join(root, '.env'), <<~STR)
            HERE_IS_A_FLOAT=3.14
            IS_IT_QUOTED="Yes it is"
            NOT_QUOTED=Hell yeah
          STR

          loon %(env)
          assert_stdout 'HERE_IS_A_FLOAT=3.14'
          assert_stdout 'IS_IT_QUOTED=Yes it is'
          assert_stdout 'NOT_QUOTED=Hell yeah'
        end
      end
    end
  end
end
